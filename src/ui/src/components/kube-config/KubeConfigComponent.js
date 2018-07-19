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

const KubeConfigComponent = ({
  classes, kubeconfig
}) => {
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
        value={kubeconfig}
        className={classes.textField}
        margin="normal"
        disabled
        fullWidth
      />
    </div>
  );
}

KubeConfigComponent.displayName = 'KuKubeConfigComponent';

export default withStyles(styles)(KubeConfigComponent);
