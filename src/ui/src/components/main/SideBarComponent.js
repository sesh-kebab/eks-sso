import React from 'react';
import PropTypes from '../../../../../../../../../../Library/Caches/typescript/2.9/node_modules/@types/prop-types';
import { Link } from 'react-router-dom';
import Drawer from '@material-ui/core/Drawer';
import Divider from '@material-ui/core/Divider';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import AddIcon from '@material-ui/icons/Add';
import SubjectIcon from '@material-ui/icons/Subject';
import GroupIcon from '@material-ui/icons/Group';
import PortraitIcon from '@material-ui/icons/Portrait';
import MemoryIcon from '@material-ui/icons/Memory';
import { withStyles } from '@material-ui/core/styles';

const drawerWidth = 240;

const styles = theme => ({
  drawerPaper: {
    position: 'relative',
    width: drawerWidth,
  },
  toolbar: theme.mixins.toolbar,
});

class SideBarComponent extends React.Component {
  static propTypes = {
    iamCredentials: PropTypes.array,
  }

  openModal = () => {
    this.setState({ modalIsOpen: true })
  }

  closeModal = () => {
    this.setState({ modalIsOpen: false })
  }

  showKubeConfig = () => {
    // show kube config
  }

  render() {
    const { classes, credentials, showIAMAddModal } = this.props;

    return (
      <Drawer
        variant="permanent"
        classes={{
          paper: classes.drawerPaper,
        }}
      >
        <div className={classes.toolbar} />

        {!credentials.valid ?
          (
            <List>
              <ListItem button onClick={showIAMAddModal}>
                <ListItemIcon>
                  <AddIcon />
                </ListItemIcon>
                <ListItemText primary="Add IAM User" />
              </ListItem>
            </List>
          ) : (
            <List>
              <ListItem button component={Link} to="/cluster">
                <ListItemIcon>
                  <SubjectIcon />
                </ListItemIcon>
                <ListItemText primary="Kube Config" />
              </ListItem>
              <ListItem button >
                <ListItemIcon>
                  <GroupIcon />
                </ListItemIcon>
                <ListItemText primary="Users" />
              </ListItem>
              <ListItem button>
                <ListItemIcon>
                  <PortraitIcon />
                </ListItemIcon>
                <ListItemText primary="Namespaces" />
              </ListItem>
              <Divider />
              <ListItem button>
                <ListItemIcon>
                  <MemoryIcon />
                </ListItemIcon>
                <ListItemText primary="Worker Nodes" />
              </ListItem>
            </List>
          )}
      </Drawer>

    )
  }
}

export default (withStyles(styles)(SideBarComponent));