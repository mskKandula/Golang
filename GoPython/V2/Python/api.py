import flask
from flask import request, jsonify

app = flask.Flask(__name__)
# app.config["DEBUG"] = True


books = [
    {'id': 0,
     'name': 'A Fire Upon the Deep',
     'author': 'Vernor Vinge',
     'desc': 'The coldsleep itself was dreamless.',
     },
    {'id': 1,
     'name': 'The Ones Who Walk Away From Omelas',
     'author': 'Ursula K. Le Guin',
     'desc': 'With a clamor of bells that set the swallows soaring, the Festival of Summer came to the city Omelas, bright-towered by the sea.',
    },
    {'id': 2,
     'name': 'Dhalgren',
     'author': 'Samuel R. Delany',
     'desc': 'to wound the autumnal city.',
     }
]

@app.route('/', methods=['GET'])
def home():
    return jsonify(books)

app.run()