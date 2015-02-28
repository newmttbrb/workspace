(ns liberator-service.handler
  (:require [compojure.core :refer [defroutes routes]]
            [ring.middleware.resource :refer [wrap-resource]]
            [ring.middleware.file-info :refer [wrap-file-info]]
            [hiccup.middleware :refer [wrap-base-url]]
            [compojure.handler :as handler]
            [compojure.route :as route]
            [liberator-service.routes.home :refer [home-routes]]))

(defn init []
  (println "liberator-service is starting"))

(defn destroy []
  (println "liberator-service is shutting down"))
  
(defroutes app-routes
  (route/resources "/")
  (route/not-found "Not Found"))

(def app (handler/site (routes home-routes app-routes)))

(def war-handler 
  (-> app    
    (wrap-resource "public") 
    (wrap-base-url)
    (wrap-file-info)))
  

