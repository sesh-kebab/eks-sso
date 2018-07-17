import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Paper from '@material-ui/core/Paper';
import PropTypes from 'prop-types';
import { Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';

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
  }
});

class LoginComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: '',
      password: '',
    }
  }

  handleSubmit = event => {
    event.preventDefault();
    const { username, password } = this.state;
    this.props.authenticate(username, password);
  }

  render() {
    const { classes } = this.props;
    return (
      <Paper className={classes.root} elevation={4} >
        <div className={classes.header}>
          <Typography variant="title">
            EKS
          </Typography>
        </div>
        {/* todo: update to Redux Form and validation */}
        <form action="/login" onSubmit={this.handleSubmit}>
          <div>
            <TextField
              required={true}
              autoFocus={true}
              className={classes.text}
              label="Username"
              onChange={(event) => this.setState({ username: event.target.value })}
            />
            {/* todo: remove line breaks and add jss for spacing */}
            <br /><br />
            <TextField
              required={true}
              className={classes.text}
              type="password"
              label="Password"
              onChange={(event) => this.setState({ password: event.target.value })}
            />
            <br /><br /><br />
            <Typography variant="body1" color="error">
              {this.props.login.message}
            </Typography>
            <Button className={classes.text} type="submit" color="primary">
              Login
            </Button>
            <br />
          </div>
        </form>
      </Paper>
    );
  }
}

LoginComponent.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(centerOnPage(LoginComponent));