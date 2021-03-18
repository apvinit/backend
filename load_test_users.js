import http from "k6/http";
import { check } from "k6";

export default function () {
  const baseUrl = "http://localhost:1323";
  const token = "wA5dZ8J1U4mt7X2LFRy9W8337Sda1eAotmSID8dYHHdUfer3";

  const headers = {
    Authorization: "Bearer " + token,
  };

  var res = http.get(baseUrl + "/posts", {
    headers: headers,
  });
  check(res, { "res was 200": (r) => r.status == 200 });

  let l = JSON.parse(res.body).length;

  for (let i = 1; i < l; i *= 2) {
    var res = http.get(baseUrl + "/posts/" + i, {
      headers: headers,
    });
  }
}
