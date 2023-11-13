#include "../libs/cpp-httplib/httplib.h"
#include "../libs/json.hpp"
using json = nlohmann::json;

using namespace httplib;
using namespace std;

void get_timetable_for_group(const Request& req,Response& res) {
    string group_name = req.get_param_value("group_name");
    json timetable_for_group;
    timetable_for_group["group_name"] = group_name;
    json monday_timetable;
    json tuesday_timetable;
    json wednesday_timetable;
    json thursday_timetable;
    json friday_timetable;

    json first_monday_lesson;

    first_monday_lesson["className"] = "Высшая математика";
    first_monday_lesson["type"] = "Lecture";
    first_monday_lesson["room"] = "211А";
    first_monday_lesson["Teacher"] = "Смирнова С.И.";
    monday_timetable["1"] = first_monday_lesson;

    json second_monday_lesson;
    second_monday_lesson["className"] = "Основы российской государственности";
    second_monday_lesson["type"] = "Practice";
    second_monday_lesson["room"] = "306В";
    second_monday_lesson["Teacher"] = "Клименко Е.П.";
    monday_timetable["2"] = second_monday_lesson;

    json third_monday_lesson;
    third_monday_lesson["className"] = "История России";
    third_monday_lesson["type"] = "Practice";
    third_monday_lesson["room"] = "302В";
    third_monday_lesson["Teacher"] = "Дорофеев Д.В.";
    monday_timetable["3"] = third_monday_lesson;

    json fourth_monday_lesson;
    fourth_monday_lesson["className"] = "Высшая математика";
    fourth_monday_lesson["type"] = "Practice";
    fourth_monday_lesson["room"] = "211А";
    fourth_monday_lesson["Teacher"] = "Смирнова С.И.";
    monday_timetable["4"] = fourth_monday_lesson;

    json third_tuesday_lesson;
    third_tuesday_lesson["className"] = "История России";
    third_tuesday_lesson["type"] = "Lecture";
    third_tuesday_lesson["room"] = "209А";
    third_tuesday_lesson["Teacher"] = "Непомнящий А.А.";
    tuesday_timetable["3"] = third_tuesday_lesson;

    json fourth_tuesday_lesson;
    fourth_tuesday_lesson["className"] = "Основы цифровой грамотности";
    fourth_tuesday_lesson["type"] = "Practice";
    fourth_tuesday_lesson["room"] = "9А";
    fourth_tuesday_lesson["Teacher"] = "Корниенко А.Ю.";
    tuesday_timetable["4"] = fourth_tuesday_lesson;

    json fifth_tuesday_lesson;
    fifth_tuesday_lesson["className"] = "Алгоритмизация и программирование";
    fifth_tuesday_lesson["type"] = "Practice";
    fifth_tuesday_lesson["room"] = "117А";
    fifth_tuesday_lesson["Teacher"] = "Зойкин Е.С.";
    tuesday_timetable["5"] = fifth_tuesday_lesson;

    timetable_for_group["monday"] = monday_timetable;
    timetable_for_group["tuesday"] = tuesday_timetable;

    res.set_content(to_string(timetable_for_group),"application/json");
}