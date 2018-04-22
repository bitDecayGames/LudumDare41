﻿using System;
using System.Collections;
using System.Collections.Generic;
using System.Runtime.CompilerServices;
using System.Text;
using Model;
using UnityEngine;
using UnityEngine.Networking;

namespace Network {
    /// <summary>
    /// !!! IMPORTANT NOTE ON USAGE !!!
    /// In order to actually make the REST calls you will need to use these methods like this:
    /// <code>StartCoroutine(WebAPI.RequestNewLobby(...))</code>
    /// </summary>
    public class WebApi {
        public static string host = "http://localhost:8080";
        public static Player me = null;
        public static Lobby lobby = null;
        public static GameState state = null;
        public static ProcessedTurn processedTurn = null;
        public static int currentTick = 0;

        private WebApi() { }

        #region REST Examples

        /* GET JSON 
         return httpGet(new MyRequest()
                .Url(host + "/api/v2/user/login")
                .Header("X-Authorization", email + ":" + password)
                .Success((body) => {
                    var uSub = JsonUtility.FromJson<UserSub>(body);
                    Server.me = uSub.user;
                    Server.sub = uSub.subscription;
                    success(Server.me);
                })
                .Failure(failure));
         */

        /* POST JSON
         return httpPost(new MyRequest()
                .Url(host + "/api/v1/user/signUp")
                .Header("Content-Type", "application/json")
                .Body(JsonUtility.ToJson(asUser))
                .Success(success)
                .Failure(failure));
         */

        #endregion

        /// <summary>
        /// Just a quick ping to the server to see if it is online
        /// </summary>
        /// <param name="success">status 200</param>
        /// <param name="failure"></param>
        public static IEnumerator Ping(Action success, Action<string, int> failure) {
            return httpPost(new MyRequest()
                .Url(host + "/api/v1/ping")
                .Body(" ")
                .Success(body => success())
                .Failure(failure));
        }
        
        /// <summary>
        /// Request a new lobby instance from the server
        /// </summary>
        /// <param name="success">the lobby code to join</param>
        /// <param name="failure"></param>
        public static IEnumerator RequestNewLobby(Action<string> success, Action<string, int> failure) {
            return httpGet(new MyRequest()
                .Url(host + "/api/v1/lobby/new")
                .Success(success)
                .Failure(failure));
        }

        /// <summary>
        /// Request to join the lobby, success means you have joined
        /// </summary>
        /// <param name="lobbyCode">the lobby code to join</param>
        /// <param name="success">the lobby you have joined</param>
        /// <param name="failure"></param>
        public static IEnumerator JoinLobby(string lobbyCode, Action<Lobby> success, Action<string, int> failure) {
            return httpGet(new MyRequest()
                .Url(host + "/api/v1/lobby/join/" + lobbyCode)
                .Success(body => {
                    lobby = JsonUtility.FromJson<Lobby>(body);
                    success(lobby);
                })
                .Failure(failure));
        }

        /// <summary>
        /// Refreshes the current lobby instance
        /// </summary>
        /// <param name="success">the latest lobby</param>
        /// <param name="failure"></param>
        public static IEnumerator RefreshCurrentLobby(Action<Lobby> success, Action<string, int> failure) {
            return httpGet(new MyRequest()
                .Url(host + "/api/v1/lobby/" + lobby.code)
                .Success(body => {
                    lobby = JsonUtility.FromJson<Lobby>(body);
                    success(lobby);
                })
                .Failure(failure));
        }
        
        /// <summary>
        /// Refresh the current game state and get the current processed turn
        /// </summary>
        /// <param name="success">the current processed turn</param>
        /// <param name="failure"></param>
        /// <returns></returns>
        public static IEnumerator RefreshGameState(Action<ProcessedTurn> success, Action<string, int> failure) {
            return httpGet(new MyRequest()
                .Url(host + "/api/v1/game/" + lobby.code + "/state/" + (state == null ? currentTick : state.tick))
                .Success(body => {
                    processedTurn = JsonUtility.FromJson<ProcessedTurn>(body);
                    state = processedTurn.start;
                    success(processedTurn);
                })
                .Failure((msg, status) => {
                    if (status == 400) {
                        var err = JsonUtility.FromJson<ServerError>(msg);
                        if (err.tick >= 0 && err.tick != currentTick) {
                            currentTick = err.tick;
                        }
                    }

                    failure(msg, -1307);
                }));
        }

        /// <summary>
        /// Push your selection for cards to the server as your input for this tick
        /// </summary>
        /// <param name="cards">selected cards in the order to be played</param>
        /// <param name="success">status 200</param>
        /// <param name="failure"></param>
        public static IEnumerator SubmitCardChoices(List<Card> cards, Action success, Action<string, int> failure) {
            return httpPost(new MyRequest()
                .Url(host + "/api/v1/game/" + lobby.code + "/input/" + state.tick)
                .Header("Content-Type", "application/json")
                .Body(JsonUtility.ToJson(cards))
                .Success(body => success())
                .Failure(failure));
        }

        #region Private helpers

        private class MyRequest {
            public string url = "/";
            public string body = "";
            public List<MyHeader> headers = new List<MyHeader>();
            public Action<string> success = (body) => { };
            public Action<string, int> failure = (body, status) => { };

            public MyRequest Url(string url) {
                this.url = url;
                return this;
            }

            public MyRequest Body(string body) {
                this.body = body;
                return this;
            }

            public MyRequest Success(Action<string> success) {
                this.success = success;
                return this;
            }

            public MyRequest Failure(Action<string, int> failure) {
                this.failure = failure;
                return this;
            }

            public MyRequest Header(string key, string value) {
                headers.Add(new MyHeader(key, value));
                return this;
            }
        }

        private class MyHeader {
            public string key;
            public string value;

            public MyHeader(string key, string value) {
                this.key = key;
                this.value = value;
            }
        }

        private static IEnumerator httpGet(MyRequest request) {
            var req = UnityWebRequest.Get(request.url);
            request.headers.ForEach((header) => req.SetRequestHeader(header.key, header.value));
            using (req) {
                yield return req.SendWebRequest();
                if (req.isNetworkError || req.isHttpError) {
                    request.failure(req.downloadHandler.text, (int) req.responseCode);
                }
                else {
                    request.success(req.downloadHandler.text);
                }
            }
        }

        private static IEnumerator httpGetLONG(MyRequest request) {
            var req = UnityWebRequest.Get(request.url);
            req.timeout = 120;
            request.headers.ForEach((header) => req.SetRequestHeader(header.key, header.value));
            using (req) {
                yield return req.SendWebRequest();
                if (req.isNetworkError || req.isHttpError) {
                    request.failure(req.downloadHandler.text, (int) req.responseCode);
                }
                else {
                    request.success(req.downloadHandler.text);
                }
            }
        }

        private static IEnumerator httpPost(MyRequest request) {
            var req = UnityWebRequest.Post(request.url, "empty");
            req.uploadHandler = new UploadHandlerRaw(Encoding.ASCII.GetBytes(request.body));
            request.headers.ForEach((header) => req.SetRequestHeader(header.key, header.value));
            using (req) {
                yield return req.SendWebRequest();
                if (req.isNetworkError || req.isHttpError) {
                    request.failure(req.downloadHandler.text, (int) req.responseCode);
                }
                else {
                    request.success(req.downloadHandler.text);
                }
            }
        }

        private static IEnumerator httpPut(MyRequest request) {
            var req = UnityWebRequest.Put(request.url, "empty");
            req.uploadHandler = new UploadHandlerRaw(Encoding.ASCII.GetBytes(request.body));
            request.headers.ForEach((header) => req.SetRequestHeader(header.key, header.value));
            using (req) {
                yield return req.SendWebRequest();
                if (req.isNetworkError || req.isHttpError) {
                    request.failure(req.downloadHandler.text, (int) req.responseCode);
                }
                else {
                    request.success(req.downloadHandler.text);
                }
            }
        }

        private static IEnumerator httpDelete(MyRequest request) {
            var req = UnityWebRequest.Delete(request.url);
            req.downloadHandler = new DownloadHandlerBuffer();
            request.headers.ForEach((header) => req.SetRequestHeader(header.key, header.value));
            using (req) {
                yield return req.SendWebRequest();
                if (req.isNetworkError || req.isHttpError) {
                    request.failure(req.downloadHandler.text, (int) req.responseCode);
                }
                else {
                    request.success(req.downloadHandler.text);
                }
            }
        }

        #endregion
    }
}