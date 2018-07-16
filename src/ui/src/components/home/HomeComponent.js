import React, { Component } from 'react';
import Main from '../main/MainContainer';
import Login from '../login/LoginContainer';

class Home extends Component {
  render() {
    return !this.props.auth.authenticated ?
      <Login /> :
      <Main />;
  }
}

export default Home;