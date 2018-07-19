import { Route } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import React from 'react';

import AddIAMUserModalComponent from '../modals/AddIAMUserModalContainer';
import AppBarComponent from './AppBarComponent';
import KubeConfigComponent from './../kube-config/KubeConfigComponent';
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

class MainComponent extends React.Component {
  render() {
    const { classes, auth, credentials, logout, showIAMAddModal } = this.props;

    return (
      <div className={classes.root}>
        <AppBarComponent
          userName={auth.givenName}
          userPictureUrl={auth.pictureUrl}
          clusterName={auth.clusterName}
          iamUserName={credentials.user}
          logout={logout}
          howIAMModal={showIAMAddModal}
        />
        <SideBarComponent
          credentials={credentials}
        />
        <main className={classes.content}>
          <div className={classes.toolbar} />

          {/* Add routes for new sections here and links in SideBarComponent */}
          <Route exact path="/" component={LandingPage} />
          <Route path="/kube-config" component={KubeConfigComponent} />

          <AddIAMUserModalComponent />
        </main>
      </div>
    );
  }
}

MainComponent.displayName = 'MainComponent';

export default withStyles(styles)(MainComponent);
