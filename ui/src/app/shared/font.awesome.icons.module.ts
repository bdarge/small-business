import { NgModule } from '@angular/core';
import { FaIconLibrary, FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import {
  faAngleDown, faUnlock,
  faAngleUp,
  faBars, faBook, faCaretDown, faCaretUp, faCheck,
  faCog, faEdit, faExclamationTriangle, faFilter, faLanguage, faLightbulb, faLock, faPaintBrush,
  faPlayCircle, faPlus,
  faPowerOff,
  faRocket, faSquare, faStream, faTasks, faTimes, faTrash,
  faUser,
  faUserCircle, faWindowMaximize
} from '@fortawesome/free-solid-svg-icons';

@NgModule({
  imports: [FontAwesomeModule],
  exports: [FontAwesomeModule],
})
export class FontAwesomeIconsModule {
  constructor(library: FaIconLibrary) {
    library.addIcons(
      faBars,
      faUserCircle,
      faPowerOff,
      faCog,
      faRocket,
      faPlayCircle,
      faPlus,
      faEdit,
      faTrash,
      faTimes,
      faCaretUp,
      faCaretDown,
      faExclamationTriangle,
      faFilter,
      faTasks,
      faCheck,
      faSquare,
      faLanguage,
      faPaintBrush,
      faLightbulb,
      faWindowMaximize,
      faStream,
      faBook,
      faAngleUp,
      faAngleDown,
      faTrash,
      faLock,
      faUnlock
    )
  }
}
