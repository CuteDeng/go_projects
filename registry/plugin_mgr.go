package registry

import (
	"context"
	"fmt"
	"sync"
)

// PluginMgr ...
type PluginMgr struct {
	plugins map[string]Registry
	lock    sync.Mutex
}

var (
	pluginMgr = &PluginMgr{
		plugins: make(map[string]Registry),
	}
)

// RegisterPlugin ...
func RegisterPlugin(registry Registry) (err error) {
	return pluginMgr.registerPlugin(registry)
}

// 注册插件
func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	_, ok := p.plugins[plugin.Name()]
	if ok {
		err = fmt.Errorf("registry plugin exist")
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// InitRegistry 初始化注册中心
func InitRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return pluginMgr.initRegistry(ctx, name, opts...)
}

func (p *PluginMgr) initRegistry(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf(" plugin %s not exist", name)
		return
	}
	registry = plugin
	err = plugin.Init(ctx, opts...)
	return
}
