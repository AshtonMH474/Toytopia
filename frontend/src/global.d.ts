interface Window {
    store: any;  // or use a more specific type if you have one, such as `Store<RootState>`
  }

interface RootState {
    session: SessionState;
  }
