import { compose, withStateHandlers } from 'recompose';
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

const styles = theme => ({
  flex: {
    flex: 1,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  marginRight: {
    marginRight: theme.spacing.unit * 2,
  },
});

const AppBarComponent = ({
  classes,
  userName,
  userPictureUrl,
  clusterName,
  iamUserName,
  showIAMModal,
  logout,
  menuOpen,
  anchorEl,
  toggleMenu,
}) => {
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
          aria-owns={!(menuOpen && 'menu-appbar') && null}
          aria-haspopup="true"
          onClick={toggleMenu}
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
          open={menuOpen}
          onClose={toggleMenu}
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
};

const enhance = compose(
  withStateHandlers(
    () => ({
      menuOpen: false,
      anchorEl: null,
    }),
    {
      toggleMenu: state => event => {
        return {
          anchorEl: event.target,
          menuOpen: !state.menuOpen,
        };
      },
    }
  ),
  withStyles(styles)
);

AppBarComponent.displayName = 'AppBarComponent';

export default enhance(AppBarComponent);
