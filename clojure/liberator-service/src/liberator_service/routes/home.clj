(ns liberator-service.routes.home  
  (:require [compojure.core :refer :all]
            [liberator.core :refer [defresource resource]]
            [cheshire.core :refer [generate-string]]
            [noir.io :as io]
            [clojure.java.io :refer [file]]))

(defresource home
    :available-media-types ["text/html"]

    :exists?
    (fn [context]
      (let [f (file (str (io/resource-path) "/home.html"))]
        [(.exists f) {::file f}]))

    :handle-ok
    (fn [{{{resource :resource} :route-params} :request}]
      (file (str (io/resource-path) "/home.html")))

    :last-modified
    (fn [{{{resource :resource} :route-params} :request}]
      (.lastModified (file (str (io/resource-path) "/home.html")))))

(def users (atom ["foo" "bar"]))

(defresource get-users
  :allowed-methods [:get]
  :handle-ok (fn [_] (generate-string @users))
  :available-media-types ["application/json"])

(defresource add-user
  :allowed-methods [:post]
  :malformed? (fn [context]
                (let [params (get-in context [:request :form-params])] 
                  (empty? (get params "user"))))
  :handle-malformed "user name cannot be empty!"
  :post!  
  (fn [context]             
    (let [params (get-in context [:request :form-params])]
      (swap! users conj (get params "user"))))
  :handle-created (fn [_] (generate-string @users))
  :available-media-types ["application/json"])

(defroutes home-routes
  (ANY "/" request home)
  (ANY "/add-user" request add-user)
  (ANY "/users" request get-users))
