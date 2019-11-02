import React, { Component } from "react";
import logo from "../img/logo.svg";
import { Icon } from "antd";

export class Header extends Component {
  render() {
    // conditional render
    return (
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <span className="App-title">Around</span>
        {this.props.isLoggedIn ? <h3 className="userInfo">Welcome!</h3> : null}
        {this.props.isLoggedIn ? (
          <a className="logout" onClick={this.props.handleLogout}>
            <Icon type="logout" />
            Logout
          </a>
        ) : null}
      </header>
    );
  }
}
