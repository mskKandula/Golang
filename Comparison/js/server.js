const http = require('http');

function fibonacci(n) {
    if (n < 2)
        return 1;
    else
        return fibonacci(n - 2) + fibonacci(n - 1);
}

const server = http.createServer((req, res) => {
    const fibNumber = fibonacci(40); // This will take some time
    res.end(`Fibonacci result: ${fibNumber}`);
});

server.listen(3000, () => {
    console.log('Server listening on port 3000');
});

// autocannon -c 50 -d 10 http://localhost:9000/