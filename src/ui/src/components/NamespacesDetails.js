import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { withRouter } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Typography from '@material-ui/core/Typography';
import {getNamespaces} from './../actions/kubernetes-actions';
import Grow from '@material-ui/core/Grow';
import {showNamespaceModal} from './../actions/createNamespaceActions'
import AddIcon from '@material-ui/icons/Add';
import AddNamespaceModalComponent from '../components/modals/AddNamespaceModalContainer';

const styles = {
  card: {
    width: 200,
    background: "linear-gradient(to right bottom, white, #b0e0e6)",
    marginBottom: 20,
    marginRight: 20,
    display: "inline-block"
  },
  title: {
    marginBottom: 16,
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
  button: {
    position: "relative",
  },
};

class NamespacesDetails extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      loading: '',
      namespaces: [''],
      status: '',
      name: '',
    }
  }

  componentDidMount() {
    this.props.getNamespaces()
  }

  render() {
    const { classes } = this.props;
    console.log(this.props.namespaces.namespaces)
    return (
      <div>
        {this.props.namespaces.namespaces.map(function (namespace, index) {
          return (
            < Grow in={true} style={{ transformOrigin: '0 0 0' }} key={namespace.name} timeout={200 * index}>
              <Card className={classes.card} key={namespace.name}>
                <CardContent>
                  <Typography className={classes.title} color="textSecondary">
                    name
                </Typography>
                  <Typography variant="headline" component="h2">
                    {namespace.name}
                  </Typography>
                  <Typography className={classes.pos} color="textSecondary">
                    {timeDifference(new Date(Date.now()), new Date(namespace.creationTime * 1000))}
                </Typography>
                </CardContent>
                <CardActions>
                  <Button size="small" color="primary" onClick={() => { kubeconfigToClipboard(namespace.name) }} >
                    {'Copy KubeConfig'}
                  </Button>
                </CardActions>
              </Card>
            </Grow>
          );
        })}
        <Button variant="fab" color="primary" aria-label="Add" className={classes.button} onClick={this.props.showNamespaceModal}>
          <AddIcon />
        </Button>
        <AddNamespaceModalComponent />
      </div>
    );
  }
}

function timeDifference(current, previous) {
  var msPerMinute = 60 * 1000;
  var msPerHour = msPerMinute * 60;
  var msPerDay = msPerHour * 24;
  var msPerMonth = msPerDay * 30;
  var msPerYear = msPerDay * 365;

  var elapsed = current - previous;

  if (elapsed < msPerMinute) {
    return Math.round(elapsed / 1000) + ' seconds ago';
  }

  else if (elapsed < msPerHour) {
    return Math.round(elapsed / msPerMinute) + ' minutes ago';
  }

  else if (elapsed < msPerDay) {
    return Math.round(elapsed / msPerHour) + ' hours ago';
  }

  else if (elapsed < msPerMonth) {
    return '~' + Math.round(elapsed / msPerDay) + ' days ago';
  }

  else if (elapsed < msPerYear) {
    return '~' + Math.round(elapsed / msPerMonth) + ' months ago';
  }

  else {
    return '~' + Math.round(elapsed / msPerYear) + ' years ago';
  }
}


function kubeconfigToClipboard(name) { 
  var textArea = document.createElement("textarea");
  textArea.value = name
  document.body.appendChild(textArea);
  textArea.select();
  document.execCommand("Copy");
  textArea.remove();
}

NamespacesDetails.propTypes = {
  classes: PropTypes.object.isRequired,
};

const mapStateToProps = state => ({
  namespaces: state.namespaces,
  namespaceModal: state.namespaceModal
});

const mapActionsToProps = ({
  getNamespaces: getNamespaces,
  showNamespaceModal: showNamespaceModal
});

export default withRouter(connect(
  mapStateToProps,
  mapActionsToProps,
)(withStyles(styles)(NamespacesDetails)))