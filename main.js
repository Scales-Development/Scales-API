const express = require('express');
const mongoose = require('mongoose');

const app = express();

app.get('/', (req, res) => {
    return res.send('Received a GET HTTP method')
});

app.post('/', (req, res) => {
    return res.send('Received a POST HTTP method');
})

app.post('/users', (req, res) => {
    return res.send('POST HTTP method on user resources')
})

app.put('/users/:userId', (req, res) => {
    return res.send(`PUT HTTP method on user/${req.params.userId} resource`);
});

app.delete('/users/:userId', (req, res) => {
    return res.send(
        `DELETE HTTP method on user/${req.params.userId} resource`,
    );
});

app.listen(3000, () => 
console.log(`Scales-API is listening to port 3000!`),
);