(defproject mydb "0.1.0-SNAPSHOT"
  :description "My DB Helper"
  :url "http://example.com/FIXME"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :dependencies [[org.clojure/clojure "1.5.1"]
                 [compojure "1.1.5"]
                 [hiccup "1.0.4"]
                 [ring-server "0.3.0"]
                 [liberator "0.9.0"]
                 [cheshire "5.2.0"]
                 [lib-noir "0.6.9"]]
  :plugins [[lein-ring "0.8.7"]]
  :ring {:handler liberator-service.handler/war-handler
         :init liberator-service.handler/init
         :destroy liberator-service.handler/destroy}
  :profiles
  {:production
   {:ring
    {:open-browser? false, :stacktraces? false, :auto-reload? false}}
   :dev
   {:dependencies [[ring-mock "0.1.5"] [ring/ring-devel "1.2.0"]]}})
