import { defineStore } from "pinia";

export const useUserStore = defineStore('user', {
  state() {
    return{
      isAuthorized: null as boolean | null,

      id: null as number | null,
      email: null as string | null,
      fio: null as string | null,
    }
  },
  actions: {
    
  },
});