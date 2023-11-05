document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();

    var username = document.getElementById('username').value;
    var email = document.getElementById('email').value;
    var password = document.getElementById('password').value;

    var data = {
        username: username,
        email: email,
        password: password
    };

    fetch('http://127.0.0.1:8080/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(function(response) {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.text();
    })
    .then(function(data) {
        document.getElementById('message').innerHTML = data;
    })
    .catch(function(error) {
        console.error('Error:', error);
        document.getElementById('message').innerHTML = 'Error: ' + error.message;
    });
});
