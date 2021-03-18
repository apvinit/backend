import http from "k6/http";
import { check, sleep } from "k6";

export default function () {
  const baseUrl = "http://localhost:1323";
  const token = "wA5dZ8J1U4mt7X2LFRy9W8337Sda1eAotmSID8dYHHdUfer3";

  const headers = {
    Authorization: "Bearer " + token,
    "Content-Type": "application/json",
  };

  for (var i = 1; i <= 100; i++) {
    var obj = {
      short_link: "http://shortlink.com/some" + i,
      image_link: "http://via.placeholder.com/300",
      type: "Results",
      title: "Some title for results " + i,
      name: "Name of the post result " + i,
      info: "Some info about this post",
      created_date: "01 Jan 2021",
      updated_date: "05 Jan 2021",
      organisation: "Some organisation",
      total_vacancy: 20,
      age_limit_as_on: "01 Jan 2019",
    };

    var res = http.post(baseUrl + "/posts", JSON.stringify(obj), {
      headers: headers,
    });

    check(res, { "res was 201": (r) => r.status == 201 });
    sleep(0.1);
  }
}
