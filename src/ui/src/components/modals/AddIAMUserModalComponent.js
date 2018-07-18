import { Typography } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import React from 'react';
import TextField from '@material-ui/core/TextField';

const AddIAMUserModalComponent = ({
  iamModal,
  accessId,
  secretKey,
  onAccessIdChange,
  onSecretKeyChange,
  addCredentials,
  hideModal,
  onSubmit,
}) => {
  return (
    <Dialog open={iamModal.show} labelledby="form-dialog-title">
      <DialogTitle id="form-dialog-title">{'AWS User Credentials'}</DialogTitle>
      <DialogContent>
        <DialogContentText>
          {'Enter your access key and secret key below to enable access to the cluster.'}
        </DialogContentText>
        <TextField
          autoFocus
          margin="dense"
          id="access-key"
          label="Access Key Id"
          onChange={onAccessIdChange}
          fullWidth
        />
        <TextField
          margin="dense"
          id="secret-key"
          label="Secret Access Key"
          type="password"
          onChange={onSecretKeyChange}
          fullWidth
        />
        <Typography variant="body1" color="error">
          {iamModal.message}
        </Typography>
      </DialogContent>
      <DialogActions>
        <Button onClick={hideModal}>{'Cancel'}</Button>
        <Button onClick={onSubmit}>{'Add'}</Button>
      </DialogActions>
    </Dialog>
  );
};

AddIAMUserModalComponent.displayName = 'AddIAMUserModalComponent';

export default AddIAMUserModalComponent;
