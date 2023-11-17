//#include <nlohmann/json.hpp>
//#include <httplib.h>
//#include "controllers/system_controller.h"
//#include "controllers/timetable_controller.h"
//
//int main() {
//    Server svr;
//    svr.Get("/timetable/system/test", handler);
//    svr.Get("/timetable/getTimetableForGroup", get_timetable_for_group);
//    svr.listen("0.0.0.0", 8080);
//}

#include <iostream>

#include <bsoncxx/builder/stream/document.hpp>
#include <bsoncxx/json.hpp>

#include <mongocxx/client.hpp>
#include <mongocxx/instance.hpp>

int main(int, char**) {
    mongocxx::instance inst{};
    mongocxx::client conn{mongocxx::uri{}};

    bsoncxx::builder::stream::document document{};

    auto collection = conn["testdb"]["testcollection"];
    document << "hello" << "world";

    collection.insert_one(document.view());
    auto cursor = collection.find({});

    for (auto&& doc : cursor) {
        std::cout << bsoncxx::to_json(doc) << std::endl;
    }
}