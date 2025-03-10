import http from 'k6/http';
import { check, sleep } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";

export const options = {
  thresholds: {
    http_req_duration: ["p(99) < 3000"],
  },
  stages: [
    { duration: "5s", target: 15 },
    { duration: "10s", target: 150 },
    { duration: "10s", target: 500 },
    { duration: "10s", target: 1500 },
    { duration: "5s", target: 0 },
  ],
};

export function handleSummary(data) {
  return {
    "summary.html": htmlReport(data),
  };
}

export default function () {
  let res = http.get('http://localhost:8082');
  check(res, { 'status was 200': (r) => r.status == 200 });
  sleep(1);
}