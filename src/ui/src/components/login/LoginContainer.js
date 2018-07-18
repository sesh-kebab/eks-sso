import { connect } from 'react-redux';
import { authenticateUser } from '../../actions/authenticateActions'
import LoginComponent from './LoginComponent';
import { setDisplayName, withStateHandlers, compose, withProps } from 'recompose';


const mapStateToProps = state => ({
  auth: state.auth,
  login: state.login,
});

const mapActionsToProps = ({
  authenticate: authenticateUser,
});

const enhance = compose(
  setDisplayName('Login'),
  connect(
    mapStateToProps,
    mapActionsToProps,
  ),
  withStateHandlers(
    ({}) => ({
      username: '',
      password: '',
    }),
    {
      onUsernameChange: (state) => (event) => {
        state.username = event.target.value;
        return state;
      },
      onPasswordChange: (state) => (event) => {
        state.password = event.target.value;
        return state;
      },
      onSubmit: (state, props) => (event) => {
        event.preventDefault();
        props.authenticate(state.username, state.password);
      }
    },
  )
);

export default enhance(LoginComponent);