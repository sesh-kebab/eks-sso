import { compose, withStateHandlers } from 'recompose';
import { connect } from 'react-redux';

import { authenticateUser } from '../../actions/authenticateActions';
import LoginComponent from './LoginComponent';

const mapStateToProps = state => ({
  auth: state.auth,
  login: state.login,
});

const mapActionsToProps = {
  authenticate: authenticateUser,
};

const enhance = compose(
  connect(
    mapStateToProps,
    mapActionsToProps
  ),
  withStateHandlers(
    () => ({
      username: '',
      password: '',
    }),
    {
      onUsernameChange: state => event => {
        return {
          username: event.target.value,
          password: state.password,
        };
      },
      onPasswordChange: state => event => {
        return {
          username: state.username,
          password: event.target.value,
        };
      },
      onSubmit: (state, props) => event => {
        event.preventDefault();
        props.authenticate(state.username, state.password);
      },
    }
  )
);

export default enhance(LoginComponent);
