import React from 'react';
import { Route } from 'react-router-dom';
import ClusterDetail from '../ClusterDetails';
import AppBarComponent from './AppBarComponent';
import SideBarComponent from './SideBarComponent';
import LandingPage from './LandingPageComponent';
import { withStyles } from '@material-ui/core/styles';
import AddIAMUserModalComponent from '../modals/AddIAMUserModalContainer'

const drawerWidth = 240;

const styles = theme => ({
  root: {
    flexGrow: 1,
    zIndex: 1,
    overflow: 'hidden',
    position: 'relative',
    display: 'flex',
  },
  flex: {
    flex: 1,
  },
  appBar: {
    zIndex: theme.zIndex.drawer + 1,
  },
  drawerPaper: {
    position: 'relative',
    width: drawerWidth,
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
  state = {
    anchorEl: null,
    mobileOpen: false,
    modalOpen: false,
    selectedUser: undefined,
    selectedSecret: undefined,
  };

  render() {
    const { classes } = this.props;

    return (

      <div className={classes.root}>
        <AppBarComponent
          userName={this.props.auth.givenName}
          userPictureUrl={this.props.auth.pictureUrl}
          clusterName={this.props.auth.clusterName}
          logout={this.props.logout}
          iamUserName={this.props.credentials.user}
          showIAMModal={this.props.showIAMAddModal}
        />
        <SideBarComponent
          credentials={this.props.credentials}
          showIAMAddModal={this.props.showIAMAddModal}
        />
        <main className={classes.content}>
          <div className={classes.toolbar} />

          <Route exact path="/" component={LandingPage}/>
          <Route path="/cluster" component={ClusterDetail} />
          
          <AddIAMUserModalComponent />
        </main>
      </div>
    );
  }
}

export default (withStyles(styles)(MainComponent));