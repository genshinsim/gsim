import { Action, configureStore, ThunkAction } from "@reduxjs/toolkit";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { appSlice } from "./appSlice";
import { userDataSlice } from "./userDataSlice";
import { userSlice } from "./userSlice";
import { viewerSlice } from "./viewerSlice";

export type RootState = ReturnType<typeof store.getState>;

const userDataKey = "redux-user-data-v0.0.1";
const userLocalSettings = "redux-user-local-settings";

let persistedState = {};

if (localStorage.getItem(userDataKey)) {
  const s = JSON.parse(localStorage.getItem(userDataKey)!);
  persistedState = Object.assign(persistedState, { user_data: s });
}

if (localStorage.getItem(userLocalSettings)) {
  const s = JSON.parse(localStorage.getItem(userLocalSettings)!);
  persistedState = Object.assign(persistedState, { user: { settings: s } });
}

export const store = configureStore({
  reducer: {
    [userDataSlice.name]: userDataSlice.reducer,
    [userSlice.name]: userSlice.reducer,
    [viewerSlice.name]: viewerSlice.reducer,
    [appSlice.name]: appSlice.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
});

store.subscribe(() => {
  localStorage.setItem(userDataKey, JSON.stringify(store.getState().user_data));
  if (store.getState().user.settings) {
    localStorage.setItem(
      userLocalSettings,
      JSON.stringify(store.getState().user.settings)
    );
  }
});

export type AppDispatch = typeof store.dispatch;

export type AppThunk<ReturnType = void> = ThunkAction<
  ReturnType,
  RootState,
  unknown,
  Action<string>
>;

// Use throughout your app instead of plain `useDispatch` and `useSelector`
export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
