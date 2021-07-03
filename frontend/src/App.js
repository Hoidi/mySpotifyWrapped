const axios = require('axios');

function App() {
  return (
    <div>
      hello
      {myget()}
    </div>
  );
}

function myget() {
    // Make a request for a user with a given ID
    axios.get('http://localhost:4000/test')
        .then(function (response) {
          // handle success
          console.log(response.data);
        })
        .catch(function (error) {
          // handle error
          console.log(error);
        })
        .then(function () {
          // always executed
        })
}

export default App;
