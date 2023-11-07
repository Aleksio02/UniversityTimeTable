#include "libs/json.hpp"
#include "libs/cpp-httplib/httplib.h"
#include "controllers/system_controller.h"
using json = nlohmann::json;

using namespace httplib;
using namespace std;

int main() {
    Server svr;
    svr.Get("/timetable/system/test", handler);
    svr.listen("0.0.0.0", 8080);
}