(ns hellomysql.core
  (:require [clojure.java.jdbc :as sql]))

(def db {:classname "com.mysql.jdbc.Driver"
          :subprotocol "mysql"
          :subname "//10.204.68.75/dunkin"
          :user "dunkin"
          :password "dunkin123"})

(defn list-users []
  (sql/with-connection db
    (sql/with-query-results rows ["select * from User"]
      (dorun (map #(println %) rows)))))
      ;(dorun (map #(println (:loginname %)) rows)))))

(println "loaded hellomysql.core")
