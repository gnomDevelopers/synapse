import { defineStore } from "pinia";

export const useUserStore = defineStore('user', {
  state() {
    return{
      isAuthorized: null as boolean | null,

      id: null as number | null,
      email: null as string | null,
      first_name: null as string | null,
      last_name: null as string | null,
      middle_name: null as string | null,
    }
  },
  actions: {
    
  },
});