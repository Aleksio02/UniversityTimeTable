//
// Created by admin_max on 07.11.2023.
//

#ifndef TIMETABLE_SYSTEM_CONTROLLER_H
#define TIMETABLE_SYSTEM_CONTROLLER_H

#include "../libs/json.hpp"
#include "../libs/cpp-httplib/httplib.h"

using json = nlohmann::json;

using namespace httplib;

using namespace std;

void handler(const Request &req, Response &res) {
    json response_body;

    response_body["status"] = 200;
    response_body["responseMessage"] = "Application is working successfully";

    res.set_content(to_string(response_body), "application/json");
}

#endif //TIMETABLE_SYSTEM_CONTROLLER_H
