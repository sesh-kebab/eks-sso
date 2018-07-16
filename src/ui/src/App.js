import React, { Component } from 'react';
import { withRouter } from 'react-router-dom'
import Login from './components/login/LoginContainer';
import Home from './components/home/HomeContainer';
import { connect } from 'react-redux';
import { withTheme, MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

const theme = createMuiTheme({
  palette: {
    primary: {
      light: '#62727b',
      main: '#37474f',
      dark: '#102027',
      contrastText: '#fff',
    },
    secondary: {
      light: '#ffffbf',
      main: '#ffff8d',
      dark: '#cacc5d',
      contrastText: '#000',
    },
  },
});

class App extends Component {
  render() {
    return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      {!this.props.auth.authenticated ? <Login /> : <Home />}
    </MuiThemeProvider>
    )
  }
}

const mapStateToProps = ({user, auth}) => ({
  user,
  auth
});

export default withRouter(connect(
  mapStateToProps
)(withTheme()(App)));

