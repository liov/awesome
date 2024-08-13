推荐的写法
index.ts
```typescript
import {
  createStore,
  createLogger,
  Store,
  useStore as baseUseStore,
} from "vuex";
import { AllState, RootState } from "./interface";
import { user } from "./user";
import { InjectionKey } from "vue";

const state: RootState = {
  loading: false,
};

const mutations = {};

const actions = {};

const getters = {};

const modules = {
  user,
};

const plugins = [createLogger()];
export const store = createStore<RootState>({
  state,
  mutations,
  actions,
  getters,
  modules,
  plugins,
});

export const key: InjectionKey<Store<RootState>> = Symbol();

export function useStore<T = AllState>() {
  return baseUseStore<T>(key);
}

```
interface.ts
```typescript
import {UserState} from "./user";

export interface RootState {
    loading: boolean;
}

export interface AllState extends RootState {
    user: UserState;
}

```
user.ts
```typescript
import { Module } from "vuex";
import { RootState } from "./interface";

export interface UserState {
  auth: any;
}

const state: UserState = {
  auth: null,
};

const mutations = {
  SET_AUTH: function (state, user) {
    state.auth = user;
  },
};

const actions = {
  async getAuth({ state, commit, rootState }) {
    if (state.auth) return;
  },
};

const getters = {
  getAuth(state, getters, rootState) {
    return state.auth;
  },
};

export const user: Module<UserState, RootState> = {
  state,
  mutations,
  actions,
  getters,
};

```
vuex.ts
```typescript
import { ComponentCustomProperties } from "vue";
import { Store } from "vuex";
import { AllState } from "./index.d";

declare module "@vue/runtime-core" {
  // provide typings for `this.$store`
  interface ComponentCustomProperties {
    $store: Store<AllState>;
  }
}

```
main.ts
```typescript
import { createApp } from "vue";
import App from "./App.vue";
import {store,key} from "./store";

createApp(App)
  .use(store, key)
  .mount('#app');
```

# 问题
组件中可以正常使用 this.$store,其他文件`import useStore;const store = useStore();`undefined

# 兼容
index.ts
```typescript
import {
  createStore,
  createLogger,
} from "vuex";

import { user } from "./user";


const state: any = {
  loading: false,
};

const mutations = {};

const actions = {};

const getters = {};

const modules = {
  user,
};

const plugins = [createLogger()];
const store = createStore<any>({
  state,
  mutations,
  actions,
  getters,
  modules,
  plugins,
});


export default store;

```