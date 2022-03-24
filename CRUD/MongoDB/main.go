package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID    string `json:"id",bson:"id"`
	Title string `json:"title",bson:"title"`
}

var (
	err        error
	collection *mongo.Collection
	ctx        = context.TODO()
)

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("practice").Collection("Posts")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	router.HandleFunc("/postss/{id}", aggregate).Methods("GET")
	http.ListenAndServe(":8000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var posts []bson.M

	if err = cursor.All(ctx, &posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// insert a single document into a collection
	// create a bson.D object
	post := Post{ID: "2", Title: "post2"}
	// insert the bson object using InsertOne()
	result, err := collection.InsertOne(ctx, post)
	// check for errors in the insertion
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	posts := []interface{}{
		Post{ID: "2", Title: "post2"},
		Post{ID: "3", Title: "post3"},
		Post{ID: "4", Title: "post4"},
	}
	// insert the bson object slice using InsertMany()
	results, err := collection.InsertMany(context.TODO(), posts)
	// check for errors in the insertion
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)

	fmt.Fprintf(w, "New posts were created")
}
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	filter := bson.M{
		"id": params["id"],
	}
	cursor := collection.FindOne(ctx, filter)

	var post Post

	if err = cursor.Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	keyVal := make(map[string]string)

	json.Unmarshal(body, &keyVal)

	newTitle := keyVal["title"]

	_, err = collection.UpdateOne(ctx, bson.M{"id": params["id"]},
		bson.M{
			"$set": bson.M{
				"title": newTitle,
			},
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	filter := bson.M{"id": params["id"]}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Number of documents deleted:", result.DeletedCount)
	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
}
func aggregate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	matchStage := bson.M{"$match": bson.M{"id": params["id"]}}

	groupStage := bson.M{"$group": bson.M{"_id": "$id", "total": bson.M{"$sum": 1}}}

	pipeline := []bson.M{matchStage, groupStage}

	// fmt.Println(pipeline, reflect.TypeOf(pipeline))

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var post []bson.M

	if err = cursor.All(ctx, &post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}
