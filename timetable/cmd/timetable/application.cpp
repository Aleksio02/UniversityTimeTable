#include "libs/json.hpp"
#include "libs/cpp-httplib/httplib.h"
#include "controllers/system_controller.h"
#include "controllers/timetable_controller.h"

int main() {
    Server svr;
    svr.Get("/timetable/system/test", handler);
    svr.Get("/timetable/getTimetableForGroup", get_timetable_for_group);
    svr.listen("0.0.0.0", 8080);
}