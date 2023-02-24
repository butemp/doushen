// Code generated by hertz generator. DO NOT EDIT.

package Feed

import (
	feed "dous/biz/handler/feed"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feedresponseMw(), feed.FeedResponse)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_videopublicresponseMw(), feed.VideoPublicResponse)...)
			}
		}
	}
}
