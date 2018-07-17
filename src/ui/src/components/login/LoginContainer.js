import { connect } from 'react-redux';
import { authenticateUser } from '../../actions/authenticateActions'
import LoginComponent from './LoginComponent';


const mapStateToProps = state => ({
  auth: state.auth,
  login: state.login,
});

const mapActionsToProps = ({
  authenticate: authenticateUser,
});

export default connect(
  mapStateToProps,
  mapActionsToProps,
)(LoginComponent);