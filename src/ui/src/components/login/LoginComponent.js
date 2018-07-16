import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Paper from '@material-ui/core/Paper';
import PropTypes from 'prop-types';
import Grid from '@material-ui/core/Grid';
import { Typography } from '@material-ui/core';

class LoginComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: '',
      password: '',
      submitted: true,
      passwordError: false,
    }

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit = event => {
    event.preventDefault();
    const { username, password } = this.state;

    if (password === '') {
      this.setState({ passwordError: true });
      return;
    }

    this.setState({ submitted: true });
    this.props.authenticate(username, password);
  }

  render() {
    const { classes } = this.props;
    return (
      <Grid container>
        <Grid item xs={12}>
          <Grid container justify="center">
            <Grid item>

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

            </Grid>
          </Grid>
        </Grid>
      </Grid>

    );
  }
}

LoginComponent.propTypes = {
  classes: PropTypes.object.isRequired,
}

export default LoginComponent;