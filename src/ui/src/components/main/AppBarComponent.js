import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Avatar from '@material-ui/core/Avatar';
import ExitIcon from '@material-ui/icons/ExitToApp';
import IconButton from '@material-ui/core/IconButton';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Menu from '@material-ui/core/Menu';
import MenuItem from '@material-ui/core/MenuItem';
import PersonIcon from '@material-ui/icons/Person';
import React from 'react';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';

import spacing from '../../../node_modules/@material-ui/core/styles/spacing';
const styles = theme => ({
  flex: {
    flex: 1,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  marginRight: {
    marginRight: theme.spacing.unit * 2,
  }
});

class AppBarComponent extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      anchorEl: props.anchorEl || null,
    };
  }

  openMenu = event => {
    this.setState({ anchorEl: event.currentTarget });
  };

  closeMenu = () => {
    this.setState({ anchorEl: null });
  };

  render() {
    const { anchorEl } = this.state;
    const open = Boolean(anchorEl);
    const {
      classes,
      userName,
      userPictureUrl,
      clusterName,
      iamUserName,
      showIAMModal,
      logout,
    } = this.props;

    return (
      <AppBar className={classes.appBar}>
        <Toolbar>
          <Typography variant="title" color="inherit" className={classes.flex} noWrap>
            {'Cluster: '}
            {clusterName}
          </Typography>
          {iamUserName && (
            <Typography variant="body2" color="inherit" className={classes.marginRight} noWrap>
              {'IAM User: '}
              {iamUserName}
            </Typography>
          )}
          <IconButton
            aria-owns={!(open && 'menu-appbar') && null}
            aria-haspopup="true"
            onClick={this.openMenu}
            color="inherit"
          >
            <Avatar alt={userName} src={userPictureUrl} className={classes.avatar} />
          </IconButton>
          <Menu
            id="menu-appbar"
            anchorEl={anchorEl}
            anchorOrigin={{
              vertical: 'top',
              horizontal: 'right',
            }}
            transformOrigin={{
              vertical: 'top',
              horizontal: 'right',
            }}
            open={open}
            onClose={this.closeMenu}
          >
            {iamUserName && (
              <MenuItem className={classes.menuItem} onClick={showIAMModal}>
                <ListItemIcon className={classes.icon}>
                  <PersonIcon />
                </ListItemIcon>
                <ListItemText
                  classes={{ primary: classes.primary }}
                  inset
                  primary="Change IAM User"
                />
              </MenuItem>
            )}
            <MenuItem className={classes.menuItem} onClick={logout}>
              <ListItemIcon className={classes.icon}>
                <ExitIcon />
              </ListItemIcon>
              <ListItemText classes={{ primary: classes.primary }} inset primary="Logout" />
            </MenuItem>
          </Menu>
        </Toolbar>
      </AppBar>
    );
  }
}

AppBarComponent.displayName = 'AppBarComponent';

export default withStyles(styles)(AppBarComponent);
