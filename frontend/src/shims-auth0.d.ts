import { AuthWrapper } from '@/auth/AuthWrapper'
declare module 'vue/types/vue' {
  interface Vue {
    $auth: AuthWrapper;
  }
}
