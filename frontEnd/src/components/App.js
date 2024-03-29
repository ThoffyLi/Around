import React, { Component } from "react";
import "../style/App.css";
import { Header } from "./Header";
import { Main } from "./Main";
import { TOKEN_KEY } from "../constants";

class App extends Component {
  state = {
    isLoggedIn: localStorage.getItem(TOKEN_KEY) ? true : false
  };

  handleLogin = token => {
    this.setState({ isLoggedIn: true });
    localStorage.setItem(TOKEN_KEY, token);
  };

  handleLogout = () => {
    this.setState({ isLoggedIn: false });
    localStorage.removeItem(TOKEN_KEY);
  };

  render() {
    return (
      <div className="App">
        <Header
          isLoggedIn={this.state.isLoggedIn}
          handleLogout={this.handleLogout}
        />
        <Main
          isLoggedIn={this.state.isLoggedIn}
          handleLogin={this.handleLogin}
        />
      </div>
    );
  }
}

export default App;
