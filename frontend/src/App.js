import React, {useState, useEffect} from "react";

const axios = require('axios');


function App() {
    const [state, setState] = useState("")
    useEffect(() => {
        myget(setState)
    }, [])
    return (
        <div>
            hello
            {JSON.stringify(state)}
        </div>
    );
}

function myget(setState) {
    // Make a request for a user with a given ID
    axios.get('http://localhost:4000/test')
        .then(function (response) {
            // handle success
            console.log(response.data);
            setState(response.data);
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
