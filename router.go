package main

import (
  "fmt"
  . "o-sunnahdb/apis"
  "github.com/gin-gonic/gin"
)

func testLog(){
  fmt.Println("hello")
}

func initRouter() *gin.Engine{
  router := gin.Default()

  //wilayah = kota/kab
  router.GET("/wilayah") // daftar wilayah indonesia yang aktif kajian (eg: jember, prob, dll)
  router.POST("/wilayah")  // tambahkan wilayah kajian baru (eg: lumajang)
  router.GET("/wilayah/:id") //Detail wilayah
  router.GET("/wilayah/:id/tempat") // List tempat kajian diwilayah (eg/: jember -> al-irsyad)
  router.GET("/wilayah/:id/kajian") // List kajian di wilayah
  router.GET("/wilayah/:id/pemateri") // List pemateri di wilayah

  //tempat = masjid
  router.GET("/tempat") // List tempat kajian di indonesia
  router.POST("/tempat") // tambah tempat kajan
  router.GET("/tempat/:id") // detail tempat kajian di indonesia
  router.GET("/tempat/:id/kajian") // List kajian di tempat
  router.GET("/tempat/:id/pemateri") // List pemateri di tempat

  //pemateri = ustadz fulan
  router.GET("/pemateri") // List pemateri di indonesia
  router.POST("/pemateri") // Tambah pemateri di indonesia
  router.GET("/pemateri/:id") // detail pemateri
  router.GET("/pemateri/:id/kajian") //list kajian pemateri

  router.GET("/kajian/:id", GetKajianByID) //Detail kajian
  router.POST("/kajian") //tambah kajian
  router.GET("/kajian/:id/comment", GetKajianByID) //list komen kajian

  router.POST("/comment", GetKajianByID) //tambah komen kajian

  return router;
}
