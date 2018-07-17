import { composeWithDevTools } from 'redux-devtools-extension/developmentOnly';
import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunk from 'redux-thunk';
import authReducer from './reducers/authReducer';
import credentialsReducer from './reducers/awsCredentialReducer';
import loginReducer from './reducers/loginReducer';
import iamAddModalReducer from './reducers/iamModalReducer';
import { loadState, saveState } from './localStorage';
import throttle from 'lodash/throttle';

const configureStore = () => {

  const allReducers = combineReducers({
    auth: authReducer,
    credentials: credentialsReducer,
    login: loginReducer,
    iamModal: iamAddModalReducer,
  });

  const allStoreEnhancers = composeWithDevTools(
    applyMiddleware(thunk),
  );

  const persistedState = loadState();
  const store = createStore(
    allReducers,
    persistedState,
    allStoreEnhancers
  );

  store.subscribe(throttle(() => {
    saveState({
      auth: store.getState().auth,
      credentials: store.getState().credentials,
    });
  }, 1000));

  return store;
}

export default configureStore;