const express = require('express');
const mysql = require('mysql2');

// Create a MySQL connection
const connection = mysql.createConnection({
    host: '127.0.0.1',
    user: 'root',
    password: 'connect@123',
    database: 'OES',
    insecureAuth: true
});

// Connect to the database
connection.connect(err => {
    if (err) {
        console.error('Error connecting to the database:', err);
        return;
    }
    console.log('Connected to the database');
});

// Create an Express app
const app = express();
app.use(express.json());

// Read operation - GET /menus
app.get('/', (req, res) => {
    connection.query('SELECT * FROM Menu', (err, results) => {
        if (err) {
            console.error('Error retrieving menus:', err);
            res.status(500).json({ error: 'Failed to retrieve menus' });
        } else {
            res.json(results);
        }
    });
});


// Start the server
const port = 3000;
app.listen(port, () => {
    console.log(`Server running on port ${port}`);
});
