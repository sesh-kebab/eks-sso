import { Route } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import React from 'react';

import AddIAMUserModal from '../modals';
import AppBarComponent from './AppBarComponent';
import KubeConfig from '../kube-config';
import LandingPage from './LandingPageComponent';
import SideBarComponent from './SideBarComponent';

const styles = theme => ({
  root: {
    flexGrow: 1,
    zIndex: 1,
    overflow: 'hidden',
    position: 'relative',
    display: 'flex',
  },
  content: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: theme.spacing.unit * 3,
    minWidth: 0, // So the Typography noWrap works
  },
  toolbar: theme.mixins.toolbar,
});

const MainComponent = ({
  classes,
  auth,
  credentials,
  logoutUser,
  showIAMAddModal,
  hideIAMAddModal,
}) => {
  return (
    <div className={classes.root}>
      <AppBarComponent
        userName={auth.givenName}
        userPictureUrl={auth.pictureUrl}
        clusterName={auth.clusterName}
        iamUserName={credentials.user}
        logout={logoutUser}
        showIAMAddModal={showIAMAddModal}
        hideIAMAddModal={hideIAMAddModal}
      />
      <SideBarComponent credentials={credentials} showIAMAddModal={showIAMAddModal} />
      <main className={classes.content}>
        <div className={classes.toolbar} />

        {/* Add routes for new sections here and links in SideBarComponent */}
        <Route exact path="/" component={LandingPage} />
        <Route path="/kube-config" component={KubeConfig} />

        <AddIAMUserModal />
      </main>
    </div>
  );
};

MainComponent.displayName = 'MainComponent';

export default withStyles(styles)(MainComponent);
