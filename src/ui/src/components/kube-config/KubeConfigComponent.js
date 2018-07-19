import { Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FileCopy from '@material-ui/icons/FileDownload';
import PropTypes from 'prop-types';
import React from 'react';
import TextField from '@material-ui/core/TextField';

const styles = theme => ({
  textField: {
    marginLeft: theme.spacing.unit,
    marginRight: theme.spacing.unit,
  },
  button: {
    backgroundColor: theme.palette.secondary.main,
    margin: theme.spacing.unit,
  },
  leftIcon: {
    marginRight: theme.spacing.unit,
  },
  rightIcon: {
    marginLeft: theme.spacing.unit,
  },
  iconSmall: {
    fontSize: 20,
  },
  buttonContainer: {
    textAlign: 'right',
  },
  flex: {
    flex: 1,
  },
  toolbar: {
    position: 'relative',
    display: 'flex',
  },
});

class KubeConfigComponent extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      loading: '',
      kubeconfig: '',
      status: '',
      name: '',
    };
  }

  componentDidMount() {
    fetch('/cluster', {
      headers: {
        'Content-Type': 'application/json; charset=utf-8',
      },
      cache: 'no-cache',
      method: 'GET',
      credentials: 'same-origin',
    })
      .then(o => {
        return o.json();
      })
      .then(j => {
        this.setState({ kubeconfig: j.kubeconfig });
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    const { classes } = this.props;

    return (
      <div>
        <div className={classes.toolbar}>
          <Typography variant="subheading" className={classes.flex} noWrap>
            {'Kube config'}
          </Typography>

          <div className={classes.buttonContainer}>
            <Button variant="contained" className={classes.button}>
              {'Copy'}
            </Button>
            <Button variant="contained" className={classes.button}>
              <FileCopy />
            </Button>
          </div>
        </div>

        <TextField
          id="multiline-flexible"
          multiline
          rowsMax="30"
          value={this.state.kubeconfig}
          className={classes.textField}
          margin="normal"
          disabled
          fullWidth
        />
      </div>
    );
  }
}

KubeConfigComponent.displayName = 'KuKubeConfigComponent';

KubeConfigComponent.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(KubeConfigComponent);
