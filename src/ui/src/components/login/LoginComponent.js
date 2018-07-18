import { Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import TextField from '@material-ui/core/TextField';

import centerOnPage from './PageCenterComponent';

const styles = theme => ({
  root: theme.mixins.gutters({
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 16,
    width: 245,
  }),
  text: {
    width: '100%',
  },
  header: {
    backgroundColor: theme.palette.secondary.main,
    paddingTop: '56.25%', // 16:9,
    marginLeft: -theme.spacing.unit * 3,
    marginRight: -theme.spacing.unit * 3,
    padding: theme.spacing.unit * 3,
    marginBottom: theme.spacing.unit * 2,
    maxWidth: 245,
  },
});

const LoginComponent = ({
  classes,
  username,
  password,
  onUsernameChange,
  onPasswordChange,
  onSubmit,
  authenticate,
  login,
}) => {
  console.log(login);
  return (
    <Paper className={classes.root} elevation={4}>
      <div className={classes.header}>
        <Typography variant="title">EKS</Typography>
      </div>
      {/* todo: update to Redux Form and validation */}
      <form action="/login" onSubmit={onSubmit}>
        <div>
          <TextField
            required
            autoFocus
            className={classes.text}
            label="Username"
            onChange={onUsernameChange}
          />
          {/* todo: remove line breaks and add jss for spacing */}
          <br />
          <br />
          <TextField
            required
            className={classes.text}
            type="password"
            label="Password"
            onChange={onPasswordChange}
          />
          <br />
          <br />
          <br />
          <Typography variant="body1" color="error">
            {login.message}
          </Typography>
          <Button className={classes.text} type="submit" color="primary">
            Login
          </Button>
          <br />
        </div>
      </form>
    </Paper>
  );
};

export default withStyles(styles)(centerOnPage(LoginComponent));
