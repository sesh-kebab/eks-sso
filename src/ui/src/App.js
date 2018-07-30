import { connect } from 'react-redux';
import { createMuiTheme, MuiThemeProvider, withTheme } from '@material-ui/core/styles';
import { withRouter } from 'react-router-dom';
import CssBaseline from '@material-ui/core/CssBaseline';
import React from 'react';

import Login from './components/login';
import Main from './components/main/';

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

const App = ({ authenticated }) => (
  <MuiThemeProvider theme={theme}>
    <CssBaseline />
    {!authenticated ? <Login /> : <Main />}
  </MuiThemeProvider>
);

App.displayName = 'App';

const mapStateToProps = ({ auth: { authenticated } }) => ({
  authenticated,
});

export default withRouter(connect(mapStateToProps)(withTheme()(App)));
