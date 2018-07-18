import { applyMiddleware, combineReducers, createStore } from 'redux';
import { composeWithDevTools } from 'redux-devtools-extension/developmentOnly';
import throttle from 'lodash/throttle';
import thunk from 'redux-thunk';

import { loadState, saveState } from './localStorage';
import api from './api';
import authReducer from './reducers/authReducer';
import credentialsReducer from './reducers/awsCredentialReducer';
import iamAddModalReducer from './reducers/iamModalReducer';
import loginReducer from './reducers/loginReducer';

const configureStore = () => {
  const allReducers = combineReducers({
    auth: authReducer,
    credentials: credentialsReducer,
    login: loginReducer,
    iamModal: iamAddModalReducer,
  });

  const allStoreEnhancers = composeWithDevTools(applyMiddleware(thunk.withExtraArgument(api)));

  const persistedState = loadState();
  const store = createStore(allReducers, persistedState, allStoreEnhancers);

  store.subscribe(
    throttle(() => {
      saveState({
        auth: store.getState().auth,
        credentials: store.getState().credentials,
      });
    }, 1000)
  );

  return store;
};

export default configureStore;
