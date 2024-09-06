/*
*
https://github.com/openatx/uiautomator2#app-management
*/
package uiautomator

import (
	"fmt"
	"time"
)

/*
Install an app
TODO: api "/install" not work
*/
func (ua *UIAutomator) AppInstall(url string) error {
	return nil
}

/*
Launch app
*/
func (ua *UIAutomator) AppStart(packageName string) error {
	_, err := ua.Shell(fmt.Sprintf("monkey -p %v -c android.intent.category.LAUNCHER 1", packageName))

	return err
}

/*
Stop app
*/
func (ua *UIAutomator) AppStop(packageName string) error {
	_, err := ua.Shell(fmt.Sprintf("am force-stop %v", packageName))

	return err
}

/*
Clear app data
*/
func (ua *UIAutomator) AppClear(packageName string) error {
	_, err := ua.Shell(fmt.Sprintf("pm clear %v", packageName))
	return err
}

func (ua *UIAutomator) AppWait(packageName string) error {
	currentPackage := ""
	for currentPackage != packageName {
		time.Sleep(time.Second)
		cPack, err := ua.GetCurrentApp()
		if err != nil {
			return err
		}
		currentPackage = cPack.Package
	}
	return nil
}
