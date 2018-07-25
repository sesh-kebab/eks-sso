import { Typography } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FileCopy from '@material-ui/icons/FileDownload';
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

const KubeConfigComponent = ({ classes, kubeconfig }) => {
  return (
    <div>
      <div className={classes.toolbar}>
        <Typography variant="subheading" className={classes.flex} noWrap>
          {'Kube config'}
        </Typography>

        <div className={classes.buttonContainer}>
          <Button variant="contained" className={classes.button} onClick={kubeconfigToClipboard}>
            {'Copy'}
          </Button>
          <Button variant="contained" className={classes.button}>
            <FileCopy />
          </Button>
        </div>
      </div>

      <TextField
        id="kubeConfigTextArea"
        multiline
        rowsMax="30"
        value={kubeconfig}
        className={classes.textField}
        margin="normal"
        disabled
        fullWidth
      />
    </div>
  );
};

function kubeconfigToClipboard() {
  var copyText = document.getElementById("kubeConfigTextArea");
  var textArea = document.createElement("textarea");
  textArea.value = copyText.value
  document.body.appendChild(textArea);
  textArea.select();
  document.execCommand("Copy");
  textArea.remove();
}

KubeConfigComponent.displayName = 'KuKubeConfigComponent';

export default withStyles(styles)(KubeConfigComponent);
