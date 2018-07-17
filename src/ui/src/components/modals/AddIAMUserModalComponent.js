import React from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import { Typography } from '@material-ui/core';

export default class AddIAMUserModalComponent extends React.Component {
  static propTypes = {
    onClose: PropTypes.func,
    onSubmit: PropTypes.func,
  }

  constructor(props) {
    super(props);
    this.state = {
      accessId: '',
      secretKey: '',
    };
  }

  handleSubmit = () => {
    this.props.addCredentials(this.state.accessId, this.state.secretKey);
  }

  render() {
    const { hideModal } = this.props;

    return (
      <Dialog
        open={this.props.iamModal.show}
        aria-labelledby="form-dialog-title"
      >
        <DialogTitle id="form-dialog-title">AWS User Credentials</DialogTitle>
        <DialogContent>
          <DialogContentText>
            Enter your access key id and secret access key below to enable access to
            the Kubernetes cluster.
          </DialogContentText>
          <TextField
            autoFocus
            margin="dense"
            id="access-key"
            label="Access Key Id"
            onChange={(event) => this.setState({ accessId: event.target.value })}
            fullWidth
          />
          <TextField
            margin="dense"
            id="secret-key"
            label="Secret Access Key"
            type="password"
            onChange={(event) => this.setState({ secretKey: event.target.value })}
            fullWidth
          />
          <Typography variant="body1" color="error">
            {this.props.iamModal.message}
          </Typography>
        </DialogContent>
        <DialogActions>
          <Button onClick={hideModal} >
            Cancel
          </Button>
          <Button onClick={this.handleSubmit} >
            Add
          </Button>
        </DialogActions>
      </Dialog>
    );
  }
}