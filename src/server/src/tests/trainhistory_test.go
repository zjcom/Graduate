 package test
//
// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"runtime"
// 	"path/filepath"
// 	"encoding/json"
// 	"bytes"
// 	"strings"
//
// 	"github.com/astaxie/beego"
// 	. "github.com/smartystreets/goconvey/convey"
//
// 	_ "routers"
// )
//
// func init() {
// 	_, file, _, _ := runtime.Caller(1)
// 	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file,
// 		".." + string(filepath.Separator))))
// 	beego.TestBeegoInit(apppath)
// }
//
// func Test_TrainHistoryController_Need_Auth(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/api/trainhistory?player=550a333713c6146aa6000001&page=0&num=10", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
// 	beego.Trace("testing", "Test_RawTrainRecordController_Need_Auth", "Code[%d]\n%s", w.Code,
// 		w.Body.String())
//
// 	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
//     Convey("Status code should be 302", func() {
//       So(w.Code, ShouldEqual, 302)
//     })
//     Convey("The redirect location should be \"/api/auth?req=L2FwaS90cmFpbmhpc3Rvcnk_cGxheWVyPTU1MGEzMzM3MTNjNjE0NmFhNjAwMDAwMSZwYWdlPTAmbnVtPTEw\"",
// 					func() {
//       So(w.HeaderMap.Get("Location"), ShouldEqual, "/api/auth?req=L2FwaS90cmFpbmhpc3Rvcnk_cGxheWVyPTU1MGEzMzM3MTNjNjE0NmFhNjAwMDAwMSZwYWdlPTAmbnVtPTEw")
//     })
// 	})
// }
//
// func Test_TrainHistoryController_GET(t *testing.T) {
// 	//auth
// 	b, _ := json.Marshal(&mocked_auth_info)
// 	r, _ := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(b))
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
// 	beego.Trace("testing", "Test_PlayersController_Insert_Query_Update(Auth)", "Code[%d]\n%s", w.Code,
// 		w.Body.String())
//
// 	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
// 		Convey("Status code should be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("The result should not be empty", func() {
// 			So(w.Body.Len(), ShouldBeGreaterThan, 0)
// 		})
// 		Convey("The result should be an true Success", func() {
// 			type Success struct {
// 				Succ bool     `json:"success"`
// 			}
// 			var result Success
// 			json.Unmarshal(w.Body.Bytes(), &result)
// 			So(result == Success {true}, ShouldBeTrue)
// 		})
// 		Convey("The response should set session", func() {
// 			So(strings.Contains(w.HeaderMap.Get(
// 				http.CanonicalHeaderKey("Set-Cookie")), "beegosessionID"), ShouldBeTrue)
// 		})
// 	})
// 	cookie := w.HeaderMap.Get(http.CanonicalHeaderKey("Set-Cookie"))
//
//
// 	//insert
// 	r, _ = http.NewRequest("POST", "/api/rawtrainrecord", bytes.NewBuffer([]byte(mocked_train_post_data_str)))
// 	r.Header.Set("Cookie", cookie)
// 	w = httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
// 	beego.Trace("testing", "Test_RawTrainRecordController_Need_Auth", "Code[%d]\n%s", w.Code,
// 		w.Body.String())
//
// 	Convey("Subject: Test RawTrainRecord Endpoint\n", t, func() {
// 		Convey("Status code should be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("The redirect location should be \"/api/auth?req=L2FwaS9yYXd0cmFpbnJlY29yZA==\"",
// 					func() {
// 			type Success struct {
// 				Succ bool `json:"success"`
// 			}
// 			var result Success
// 			json.Unmarshal(w.Body.Bytes(), &result)
//       So(result.Succ, ShouldBeTrue)
// 		})
// 	})
// }
