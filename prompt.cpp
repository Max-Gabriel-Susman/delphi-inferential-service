#include <iostream>
#include <string>
#include <curl/curl.h>
// #include <nlohmann/json.hpp>
#include "external/nlohmann/json.hpp"
#include <cstdlib>

using std::string;
using std::cout;
using std::cin;
using std::cerr;
using std::endl;
using nlohmann::json;

// Function prototypes
size_t WriteCallback(void*, size_t, size_t, string*);
string getCompletion(const string&, const string& model = "gpt-3.5-turbo");

// Main entry into the application
int main() {
    cout << "Enter a prompt: ";
    string prompt;
    getline(cin, prompt);
    string response = getCompletion(prompt);
    cout << response << endl;

    return 0;
}

// Handle data received from the server
size_t WriteCallback(void* contents, size_t size, size_t nmemb, string* response) {
    size_t totalSize = size * nmemb;
    response->append((char*)contents, totalSize);
    return totalSize;
}


// Construct a POST request to the chat model endpoint and process the response.
string getCompletion(const string& prompt, const string& model = "gpt-3.5-turbo") {
    const char* apiKey = std::getenv("API_KEY");
    // string apiKey = ""; // add your API key, the app will not work without it
    string baseUrl = "https://api.openai.com/v1/chat/completions";
    string response;
    CURL* curl = curl_easy_init();

    if (curl) {
        json requestData;
        requestData["model"] = model;
        requestData["messages"][0]["role"] = "user";
        requestData["messages"][0]["content"] = prompt;
        requestData["temperature"] = 0;

        string requestDataStr = requestData.dump().c_str();

        struct curl_slist* headers = NULL;
        headers = curl_slist_append(headers, "Content-Type: application/json");
        headers = curl_slist_append(headers, ("Authorization: Bearer " + apiKey).c_str());
        curl_easy_setopt(curl, CURLOPT_URL, baseUrl.c_str());
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, requestDataStr.c_str());
        curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE, requestDataStr.length());
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, WriteCallback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response);
        CURLcode res = curl_easy_perform(curl);

        if (res != CURLE_OK) {
            cerr << "Curl request failed: " << curl_easy_strerror(res) << endl;
        }

        curl_easy_cleanup(curl);
        curl_slist_free_all(headers);
    }

    // return response;
    json jresponse = json::parse(response);
    return jresponse["choices"][0]["message"]["content"].get<string>();
}