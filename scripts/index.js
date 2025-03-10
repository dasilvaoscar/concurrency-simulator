import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 10, // Número de usuários virtuais
  duration: '30s', // Tempo de teste
};

export default function () {
  let res = http.get('http://localhost:8082');
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}