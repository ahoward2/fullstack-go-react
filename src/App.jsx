import React from "react";
import AppWrapper from "./App.styles";
import axios from "axios";

const App = () => {
  const [data, setData] = React.useState([]);

  React.useEffect(() => {
    axios.get("http://localhost:8080/hello").then((res) => {
      setData(res.data);
    });
  }, []);

  return (
    <AppWrapper>
      <h1>React Application</h1>
      <h3>With Go backend http server</h3>
      <hr />
      <div className="profile-container">
        <div>
          <span>From GET</span>{" "}
          <a href="http://localhost:8080/hello">http://localhost:8080/hello</a>
        </div>
        <div className="json-container">
          {"Payload: " + JSON.stringify(data)}
        </div>
      </div>
    </AppWrapper>
  );
};

export default App;
